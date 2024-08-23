import { StrictMode } from "react";
import { createRoot } from "react-dom/client";
import App from "./App.tsx";
import "./index.css";
import { AuthProvider } from "./components/AuthProvider.tsx";
import { QueryClient, QueryClientProvider } from "@tanstack/react-query";

const queryClient = new QueryClient();

createRoot(document.getElementById("root")!).render(
	<StrictMode>
		<AuthProvider>
			<QueryClientProvider client={queryClient}>
				<App />
			</QueryClientProvider>
		</AuthProvider>
	</StrictMode>
);
