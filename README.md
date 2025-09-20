# Mortgage Underwriting Application

A full-stack mortgage underwriting web application with a Go backend and a React (Vite) frontend.

The calculations for the decision are set by usual real-life rules searched in Google. Some of this are:

- Property value and loan amount not under $75.000 and $50.000, respectivelly
- A special rule for high credit score, allowing up to 45% dti and 95% ltv
- Minimum credit score of 620 for primary occupancy and 680 for secondary/investment.
- Maximum of 90% ltv for primary occupancy and 80% for seconday/investment.
- Limit maximum dti to 36% when credit score below 700.
- Not more than 50% dti in any case
- Default approval when score beyond 700, dti at most 43% and ltv of 90% or lower for primary occupancy or 80% for the secondary/investment.

---

## Features

- 🏦 Mortgage decision engine with realistic business rules
- 🚀 Modern React frontend (Vite, TypeScript, TailwindCSS)
- 🔒 Go backend with REST API and PostgreSQL support
- 📦 Dockerized for easy deployment

---

## Project Structure

```
mortgage_underwriting/
├── backend/      # Go API server
├── frontend/     # React web app
```

---

## Getting Started

### Prerequisites

- [Docker](https://www.docker.com/) (recommended)
- Or: Go 1.21+, Node.js 20+ (for local dev)

---

## Development

### Backend

Remember to set your DB connection string in the a .env file. Copy the .env.example and replace with the desired value

```sh
cd backend
go run main.go
```
- The backend will start on `http://localhost:8080`.

### Frontend

```sh
cd frontend
npm install
npm run dev
```
- The frontend will start on `http://localhost:5173`.

---

## API Endpoints

- `POST /api/request-loan` — Submit a loan application
- `GET /api/loan-history` — List all loan decisions

---

## Docker Usage

### Build and Run Backend

```sh
docker build -t mortgage-backend ./backend
docker run -p 8080:8080 mortgage-backend
```

### Build and Run Frontend

```sh
docker build -t mortgage-frontend ./frontend
docker run -p 3000:80 mortgage-frontend
```

- Frontend will be available at `http://localhost:3000`
- Backend will be available at `http://localhost:8080`

---

## Docker Compose (Optional)

Create a `docker-compose.yml` to run both services together:

```yaml
version: "3.8"
services:
  backend:
    build: ./backend
    ports:
      - "8080:8080"
  frontend:
    build: ./frontend
    ports:
      - "3000:80"
```

Run both:
```sh
docker-compose up --build
```

---

## Customization

- **Business rules:** See `backend/functions/loans.go`
- **Frontend:** See `frontend/app/welcome/welcome.tsx`

---

## License

MIT

---

Built with ❤️ by Alberto Pumarejo
