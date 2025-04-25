package com.example.gateway.config;

import org.springframework.stereotype.Component;

import java.util.HashMap;
import java.util.Map;
import java.util.Set;

@Component
public class RoleProtectedRoutes {

    private final Map<String, Set<String>> protectedRoutes = new HashMap<>();

    public RoleProtectedRoutes() {
        protectedRoutes.put("POST /api/v1/admin/products", Set.of("ADMIN"));
        protectedRoutes.put("DELETE /api/v1/admin/products", Set.of("ADMIN"));
        protectedRoutes.put("PUT /api/v1/admin/products", Set.of("ADMIN"));
    }

    public Set<String> getAllowedRoles(String method, String path) {
        return protectedRoutes.getOrDefault(method + " " + path, Set.of());
    }

    public boolean isProtected(String method, String path) {
        return protectedRoutes.containsKey(method + " " + path);
    }
}

