import {
	Navigate,
	Route,
	BrowserRouter as Router,
	Routes,
} from "react-router-dom";
import AddListingForm from "./components/AddListingForm";
import ListingDisplay from "./components/ListingDisplay";
import LoginForm from "./components/LoginForm";
import Sidebar from "./components/Sidebar";
import { useAuth } from "./hooks/useAuth";
import Logout from "./components/Logout";

const ProtectedRoute = ({ children }: { children: React.ReactNode }) => {
	const { isAuthenticated } = useAuth();
	return isAuthenticated ? <>{children}</> : <Navigate to="/login" />;
};

function App() {
	const { isAuthenticated } = useAuth();

	return (
		<Router>
			<div className="flex">
				{isAuthenticated && <Sidebar />}
				<Routes>
					<Route path="/" element={<h1>HompimRent</h1>} />
					<Route path="/login" element={<LoginForm />} />
					<Route
						path="/add"
						element={
							<ProtectedRoute>
								<AddListingForm />
							</ProtectedRoute>
						}
					/>
					<Route path="/listings" element={<ListingDisplay />} />
					<Route path="/logout" element={<Logout />} />
				</Routes>
			</div>
		</Router>
	);
}

export default App;
