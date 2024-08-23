import { SubmitHandler, useForm } from "react-hook-form";
import { zodResolver } from "@hookform/resolvers/zod";
import { z } from "zod";
import { LoginCredentials } from "../types/interfaces";
import { useAuth } from "../hooks/useAuth";
import { useNavigate } from "react-router-dom";

const loginSchema = z.object({
	email: z
		.string()
		.email("Invalid email address")
		.nonempty("Email is required"),
	password: z.string().min(6, "Password must be at least 6 chracters long"),
});

type LoginFormInputs = z.infer<typeof loginSchema>;

function LoginForm() {
	const { login } = useAuth();
	const navigate = useNavigate();
	const {
		register,
		handleSubmit,
		formState: { errors },
	} = useForm<LoginFormInputs>({
		resolver: zodResolver(loginSchema),
	});

	const onSubmit: SubmitHandler<LoginFormInputs> = async (data) => {
		try {
			await login(data as LoginCredentials);
			console.log("Login successful");
			navigate("/");
		} catch (error) {
			if (error instanceof Error) {
				console.error("Login failed:", error.message);
			} else {
				console.error("Login failed:", "An unknown error occurred");
			}
		}
	};

	return (
		<form
			onSubmit={handleSubmit(onSubmit)}
			className="max-w-sm mx-auto my-8 p-4 border rounded shadow-md"
		>
			<div className="mb-4">
				<label
					htmlFor="email"
					className="block text-sm font-medium mb-1"
				>
					Email
				</label>
				<input
					id="email"
					type="email"
					{...register("email")}
					className="w-full px-3 py-2 border rounded"
				/>
				{errors.email && (
					<p className="text-red-500 text-sm">
						{errors.email.message}
					</p>
				)}
			</div>
			<div className="mb-4">
				<label
					htmlFor="password"
					className="block text-sm font-medium mb-1"
				>
					Password
				</label>
				<input
					id="password"
					type="password"
					{...register("password")}
					className="w-full px-3 py-2 border rounded"
				/>
				{errors.password && (
					<p className="text-red-500 text-sm">
						{errors.password.message}
					</p>
				)}
			</div>
			<button
				type="submit"
				className="w-full py-2 px-4 bg-blue-500 text-white rounded hover:bg-blue-600"
			>
				Login
			</button>
		</form>
	);
}

export default LoginForm;
