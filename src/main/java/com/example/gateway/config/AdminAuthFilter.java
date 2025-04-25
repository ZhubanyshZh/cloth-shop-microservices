package com.example.gateway.config;

import io.jsonwebtoken.*;
import lombok.RequiredArgsConstructor;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.beans.factory.annotation.Value;
import org.springframework.cloud.gateway.filter.GatewayFilter;
import org.springframework.cloud.gateway.filter.GatewayFilterChain;
import org.springframework.core.io.buffer.DataBuffer;
import org.springframework.http.HttpHeaders;
import org.springframework.http.HttpStatus;
import org.springframework.http.MediaType;
import org.springframework.stereotype.Component;
import org.springframework.web.server.ServerWebExchange;
import reactor.core.publisher.Mono;

import java.nio.charset.StandardCharsets;
import java.util.Set;

@Component
@RequiredArgsConstructor
public class AdminAuthFilter implements GatewayFilter {

    @Value("${jwt.secret}")
    private String SECRET_KEY;

    private final RoleProtectedRoutes roleProtectedRoutes;

    @Override
    public Mono<Void> filter(ServerWebExchange exchange, GatewayFilterChain chain) {
        String method = exchange.getRequest().getMethod().name();
        String path = exchange.getRequest().getPath().toString();

        if (!roleProtectedRoutes.isProtected(method, path)) {
            return chain.filter(exchange);
        }

        String authHeader = exchange.getRequest().getHeaders().getFirst(HttpHeaders.AUTHORIZATION);
        if (authHeader == null || !authHeader.startsWith("Bearer ")) {
            return this.onError(exchange, "Missing or invalid Authorization header", HttpStatus.UNAUTHORIZED);
        }

        String token = authHeader.substring(7);
        try {
            Claims claims = Jwts.parser()
                    .setSigningKey(SECRET_KEY.getBytes(StandardCharsets.UTF_8))
                    .parseClaimsJws(token)
                    .getBody();

            String role = claims.get("user_role", String.class);
            Set<String> allowedRoles = roleProtectedRoutes.getAllowedRoles(method, path);

            if (!allowedRoles.contains(role)) {
                return this.onError(exchange, "Access denied for role: " + role, HttpStatus.FORBIDDEN);
            }

        } catch (JwtException e) {
            return this.onError(exchange, "Invalid JWT token", HttpStatus.UNAUTHORIZED);
        }

        return chain.filter(exchange);
    }

    private Mono<Void> onError(ServerWebExchange exchange, String message, HttpStatus status) {
        exchange.getResponse().setStatusCode(status);
        exchange.getResponse().getHeaders().setContentType(MediaType.APPLICATION_JSON);
        DataBuffer buffer = exchange.getResponse().bufferFactory().wrap(("{\"error\":\"" + message + "\"}").getBytes());
        return exchange.getResponse().writeWith(Mono.just(buffer));
    }
}

