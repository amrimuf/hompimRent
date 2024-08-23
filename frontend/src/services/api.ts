import axios from "axios";

const API_URL = import.meta.env.VITE_API_URL;

const api = axios.create({
	baseURL: API_URL,
});

api.interceptors.request.use(
	(config) => {
		if (config.url && config.url.startsWith("/protected")) {
			const token = localStorage.getItem("jwtToken");
			if (token && config.headers) {
				config.headers.Authorization = `Bearer ${token}`;
			}
		}
		return config;
	},
	(error) => Promise.reject(error)
);

export default api;
