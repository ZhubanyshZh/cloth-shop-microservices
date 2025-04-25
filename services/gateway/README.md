# 🛍️ E-Commerce Platform — Microservices Architecture

Welcome to the monorepo for our large-scale **e-commerce platform** specializing in **men's clothing**. The platform is designed using a **polyglot microservices architecture** with scalability, security, and maintainability in mind.

---

## 🚀 Stack Overview

| Layer         | Tech                                      |
|--------------|------------------------------------------|
| **Frontend**  | React, Redux, Axios, PWA support       |
| **API Gateway** | Java + Spring Cloud Gateway          |
| **Core Services** | Go (Gin), Node.js (Express), Java (Spring Boot) |
| **Databases** | PostgreSQL, MongoDB, Redis, ClickHouse |
| **Messaging** | Kafka, gRPC, REST                      |
| **DevOps**    | Docker, Kubernetes, GitHub Actions, ArgoCD |
| **Monitoring** | Prometheus, Grafana, Loki, ELK         |

---

## 🧩 Services Breakdown

| Service                 | Tech      | Description                              |
|-------------------------|----------|------------------------------------------|
| **auth-service**         | Java     | JWT, OAuth2, Google, Keycloak           |
| **payment-service**      | Java     | KaspiPay, Stripe, PayPal integration    |
| **api-gateway**         | Java     | Traffic routing, security, rate limiting |
| **product-service**      | Go       | Products, images, filters, Cloudinary   |
| **inventory-service**    | Go       | Stock tracking, warehouses              |
| **cart-service**        | Go       | Shopping cart, sync with products       |
| **order-service**       | Go       | Order creation, status tracking         |
| **shipping-service**     | Go       | Delivery API integrations               |
| **notification-service** | Go       | Email, SMS, Telegram bot                |
| **analytics-service**    | Go       | Sales data, ClickHouse, reporting       |
| **recommendation-service** | Go    | Personalization engine                  |
| **promo-service**       | Node.js  | Coupons, promo codes                    |
| **review-service**      | Node.js  | Reviews, moderation                      |
| **user-service**        | Node.js  | Profiles, addresses                     |
| **admin-api**          | Node.js  | Admin panel, management                 |
| **cms-service**        | Node.js  | Content, blog, SEO pages                |

---

## 🔐 Security

- OAuth2 via Google & Keycloak
- Role-based access (Admin/User/Vendor)
- API Gateway protection (Spring Security + JWT)
- Audit Logging per service

---

## 📈 Monitoring & Observability

- **Prometheus + Grafana** for metrics
- **Loki + Grafana** for logs
- **Sentry** for frontend/backend error tracing
- **Kafka + ClickHouse** for event-based analytics

---

## 📦 CI/CD

- **GitHub Actions** for tests and Docker builds
- **ArgoCD** for Kubernetes GitOps deployments
- **Helm charts** for service configuration

---

## 🌐 External Integrations

| Service    | Integration |
|------------|--------------------------------|
| **Auth**   | Google OAuth / Keycloak       |
| **Payments** | Stripe, PayPal, Kaspi Pay   |
| **Image Hosting** | Cloudinary / S3        |
| **Notifications** | SendGrid, Twilio, Telegram |
| **Shipping APIs** | KazPost, Boxberry      |
| **Search** | ElasticSearch                 |
| **Caching** | Redis / Caffeine             |

---

## 🚀 Deployment

To run the project using Docker Compose:
```sh
docker-compose up --build
```

For Kubernetes deployment:
```sh
kubectl apply -f k8s/
```

---

## 📞 Contact

For questions or contributions, contact **Zhubanysh** at **zubanyszarylkasyn@gmail.com, zhubanysh.zharylkassynov@narxoz.kz** or visit **ZhubanyshZh**.

