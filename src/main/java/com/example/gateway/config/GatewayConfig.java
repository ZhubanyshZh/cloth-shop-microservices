package com.example.gateway.config;

import lombok.RequiredArgsConstructor;
import org.springframework.cloud.gateway.route.RouteLocator;
import org.springframework.cloud.gateway.route.builder.RouteLocatorBuilder;
import org.springframework.context.annotation.Bean;
import org.springframework.context.annotation.Configuration;

@Configuration
@RequiredArgsConstructor
public class GatewayConfig {

    private final AdminAuthFilter adminFilter;

    @Bean
    public RouteLocator customRouteLocator(RouteLocatorBuilder builder) {
        return builder.routes()
                .route("product-public", r -> r.path("/api/v1/products/**")
                        .uri("http://localhost:8080"))
                .route("product-private", r -> r.path("/api/v1/admin/products/**")
                        .filters(f -> f.filter(adminFilter))
                        .uri("http://localhost:8080"))
                .route("auth-service", r -> r.path("/api/v1/auth/**")
                        .uri("http://localhost:8081"))
                .build();
    }
}

