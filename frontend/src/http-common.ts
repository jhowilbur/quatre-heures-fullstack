import axios, { AxiosInstance } from "axios";

// TODO fix error with dotenv
// import dotenv from "dotenv";
// dotenv.config();

const apiClient: AxiosInstance = axios.create({
  // baseURL: process.env.BACKEND_ADDR, // TODO fix error with dotenv
  baseURL: "http://localhost:8081",
  headers: {
    "Content-type": "application/json",
  },
});

export default apiClient;
