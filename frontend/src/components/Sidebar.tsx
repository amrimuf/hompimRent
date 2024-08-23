import { useState } from "react";
import { useNavigate } from "react-router-dom";
import { useAuth } from "../hooks/useAuth";
import { MenuItem } from "../types/interfaces";
import { IoLogOut, IoSettingsSharp } from "react-icons/io5";
import { FaHome, FaBars } from "react-icons/fa";

const Sidebar = () => {
	const [collapsed, setCollapsed] = useState(false);
	const { logout } = useAuth();
	const navigate = useNavigate();

	const handleLogout = () => {
		logout();
		navigate("/login");
	};

	const menuItems: MenuItem[] = [
		{
			label: "Home",
			icon: <FaHome />,
			href: "/",
		},
		{
			label: "Listings",
			icon: <FaHome />,
			href: "/listings",
		},
		{
			label: "Rentals",
			icon: <FaHome />,
			href: "/rentals",
		},
		{
			label: "Settings",
			icon: <IoSettingsSharp />,
			href: "/settings",
		},
		{
			label: "Logout",
			icon: <IoLogOut />,
			onClick: handleLogout,
		},
	];

	return (
		<div
			className={`
  bg-white dark:bg-gray-800 
  h-screen 
  transition-all duration-300 ease-in-out
  shadow-lg
  ${collapsed ? "w-20" : "w-64"}
  flex flex-col
`}
		>
			<button
				onClick={() => setCollapsed(!collapsed)}
				className="p-4 hover:bg-gray-100 dark:hover:bg-gray-700 transition-colors duration-200"
			>
				<FaBars className="text-gray-500 dark:text-gray-400" />
			</button>

			<nav className="flex-grow">
				{menuItems.map((item, index) => (
					<a
						key={index}
						href={item.href}
						onClick={
							item.label === "Logout" ? handleLogout : undefined
						}
						className={`
          flex items-center
          p-4 
          hover:bg-gray-100 dark:hover:bg-gray-700 
          transition-colors duration-200
          cursor-pointer
        `}
					>
						<div
							className={`text-gray-500 dark:text-gray-400 ${
								collapsed ? "mx-auto" : "mr-4"
							}`}
						>
							{item.icon}
						</div>
						{!collapsed && (
							<span className="text-gray-700 dark:text-gray-200">
								{item.label}
							</span>
						)}
					</a>
				))}
			</nav>
		</div>
	);
};

export default Sidebar;
