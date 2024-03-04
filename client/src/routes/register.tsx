import { createSignal, useContext } from "solid-js";
import { AuthContext } from "../context/auth.context";
import { A, useNavigate } from "@solidjs/router";
import InputGroup from "../components/InputGroup";
import PasswordInputGroup from "../components/PasswordInputGroup";

function Register() {
	const [isPasswordVisible, setIsPasswordVisible] = createSignal(false);

	let formData = {
		name: "",
		email: "",
		password: "",
		confirmPassword: "",
		loginDirectly: true,
	};

	const [errors, setErrors] = createSignal<{ field: string; value: string }>();

	const Auth = useContext(AuthContext);

	const navigate = useNavigate();

	const handleSubmit = async (e: Event) => {
		e.preventDefault();

		if (formData.password !== formData.confirmPassword) {
			setErrors({
				field: "confirmPassword",
				value: "Passwords don't match",
			});
			return;
		}

		const res = await (
			await fetch("http://localhost:8080/register", {
				method: "POST",
				body: JSON.stringify({
					name: formData.name,
					email: formData.email,
					password: formData.password,
					loginDirectly: formData.loginDirectly,
				}),
				credentials: "include",
			})
		).json();
		if (res.message === "success") {
			if (formData.loginDirectly) {
				Auth?.refetch();
				navigate("/", { replace: true });
				return;
			}
			navigate("/login", { replace: true });
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
					<p class="text-center text-xl">Create a new account</p>
					<p class="text-center text-sm">
						or{" "}
						<A
							href={"/login"}
							class="text-fuchsia-600 dark:text-amber-500 hover:underline"
						>
							login
						</A>{" "}
						if you already have one
					</p>
				</div>
				<InputGroup
					inputAttrs={{
						type: "text",
						oninput: (e) => {
							setErrors(undefined);
							formData.name = e.target.value;
						},
						placeholder: "Name",
					}}
					errors={errors}
					fieldName="name"
				/>
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
				<InputGroup
					inputAttrs={{
						type: isPasswordVisible() ? "text" : "password",
						oninput: (e) => {
							setErrors(undefined);
							formData.confirmPassword = e.target.value;
						},
						placeholder: "Confirm password",
					}}
					errors={errors}
					fieldName="confirmPassword"
				/>
				<label
					class="flex items-center justify-between cursor-pointer px-3"
					for="loginDirectly"
				>
					<span>Login directly?</span>
					<input
						type="checkbox"
						class="sr-only peer"
						id="loginDirectly"
						onChange={(e) => (formData.loginDirectly = e.target.checked)}
						checked={formData.loginDirectly}
					/>
					<div class="relative w-11 h-6 bg-slate-200 peer-focus:outline-none peer-focus:ring-2 peer-focus:ring-blue-300 dark:peer-focus:ring-blue-400 rounded-full peer dark:bg-zinc-950 peer-checked:after:translate-x-full rtl:peer-checked:after:-translate-x-full peer-checked:after:border-white after:content-[''] after:absolute after:top-[2px] after:start-[2px] after:bg-white after:border-slate-300 after:border after:rounded-full after:h-5 after:w-5 after:transition-all dark:border-slate-600 peer-checked:bg-blue-500"></div>
				</label>
				<button class="outline-none border-none w-full dark:bg-indigo-700 bg-slate-900 text-white p-3 rounded-full focus-visible:ring-4 ring-blue-400/80 font-bold">
					Submit
				</button>
			</form>
		</main>
	);
}

export default Register;
