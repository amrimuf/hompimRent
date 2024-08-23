import { LoginCredentials, RegisterData } from "../types/interfaces";
import api from "./api";

export const login = async (credentials: LoginCredentials) => {
	const response = await api.post("/auth/login", credentials);
	return response.data;
};

export const signup = async (userData: RegisterData) => {
	const response = await api.post("/auth/register", userData);
	return response.data;
};
