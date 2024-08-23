import axios from "axios";
import { LoginCredentials, RegisterData } from "../../types/interfaces";

const API_URL = import.meta.env.VITE_API_URL;

export const login = async (credentials: LoginCredentials) => {
	const response = await axios.post(`${API_URL}/auth/login`, credentials);
	return response.data;
};

export const signup = async (userData: RegisterData) => {
	const response = await axios.post(`${API_URL}/auth/register`, userData);
	return response.data;
};
