import { createSignal, useContext } from "solid-js";
import { AuthContext } from "../context/auth.context";
import { A, useNavigate } from "@solidjs/router";
import { Response, TError } from "../types";
import InputGroup from "../components/InputGroup";
import PasswordInputGroup from "../components/PasswordInputGroup";

function Login() {
	const [isPasswordVisible, setIsPasswordVisible] = createSignal(false);

	let formData = {
		email: "",
		password: "",
	};

	const [errors, setErrors] = createSignal<TError>();

	const Auth = useContext(AuthContext);

	const navigate = useNavigate();

	const handleSubmit = async (e: Event) => {
		e.preventDefault();

		const res: Response<TError | undefined> = await (
			await fetch("http://localhost:8080/login", {
				method: "POST",
				body: JSON.stringify(formData),
				credentials: "include",
			})
		).json();

		if (res.message === "success") {
			Auth?.refetch();
			navigate("/", { replace: true });
			return;
		}

		setErrors(res.data.value);
	};

	return (
		<main class="pt-20 p-8 lg:px-[15dvw]">
			<form
				class="mx-auto w-[34rem] ring-slate-700 dark:ring-zinc-700 ring-2 rounded-3xl p-12 space-y-10"
				onsubmit={handleSubmit}
			>
				<div>
					<p class="text-center text-xl">Login to your account</p>
					<p class="text-center text-sm">
						or{" "}
						<A
							href={"/register"}
							class="text-fuchsia-600 dark:text-amber-500 hover:underline"
						>
							register
						</A>{" "}
						for a new account
					</p>
				</div>
				<InputGroup
					inputAttrs={{
						type: "email",
						oninput: (e) => {
							setErrors(undefined);
							formData.email = e.target.value;
						},
						placeholder: "Email",
					}}
					errors={errors}
					fieldName="email"
				/>

				<PasswordInputGroup
					inputAttrs={{
						oninput: (e) => {
							setErrors(undefined);
							formData.password = e.target.value;
						},
						placeholder: "Password",
					}}
					errors={errors}
					isPasswordVisible={isPasswordVisible}
					setIsPasswordVisible={setIsPasswordVisible}
				/>
				<button class="outline-none border-none w-full dark:bg-indigo-700 bg-slate-900 text-white p-3 rounded-full focus-visible:ring-4 ring-blue-400/80 font-bold">
					Submit
				</button>
			</form>
		</main>
	);
}

export default Login;
