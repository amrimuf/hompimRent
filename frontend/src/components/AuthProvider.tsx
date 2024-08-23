import { createContext, ReactNode, useState } from "react";
import { AuthContextType, LoginCredentials } from "../types/interfaces";
import { login as authLogin } from "./Auth/authService";

export const AuthContext = createContext<AuthContextType | undefined>(
	undefined
);
export const AuthProvider = ({ children }: { children: ReactNode }) => {
	const [isAuthenticated, setIsAuthenticated] = useState<boolean>(() => {
		const token = localStorage.getItem("jwtToken");
		return !!token;
	});


	const login = async (credentials: LoginCredentials) => {
		try {
			const result = await authLogin(credentials);
			const token = result.token;
			localStorage.setItem("jwtToken", token);
			setIsAuthenticated(true);
		} catch (error) {
			console.log("Login error", error);
		}
	};
	const logout = () => {
		localStorage.removeItem("jwtToken");
		setIsAuthenticated(false);
	};

	return (
		<AuthContext.Provider value={{ isAuthenticated, login, logout }}>
			{children}
		</AuthContext.Provider>
	);
};
