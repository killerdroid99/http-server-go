import { createContext, createResource, ParentProps, Resource } from "solid-js";
import { TAuthData } from "../types";

interface Response<T> {
	status: number;
	message: "success" | "error";
	data: {
		value: T;
	};
}

const fetchMe = async (): Promise<Response<TAuthData>> => {
	const res = await fetch("http://localhost:8080/me", {
		credentials: "include",
	});

	return res.json();
};

export const AuthContext = createContext<{
	data: Resource<Response<TAuthData>>;
	refetch: (info?: unknown) =>
		| Response<{
				userId: string;
				userName: string;
		  }>
		| Promise<
				| Response<{
						userId: string;
						userName: string;
				  }>
				| undefined
		  >
		| null
		| undefined;
}>();

export function AuthContextProvider(props: ParentProps) {
	const [data, { refetch }] = createResource(fetchMe);

	return (
		<AuthContext.Provider
			value={{
				data,
				refetch,
			}}
		>
			{props.children}
		</AuthContext.Provider>
	);
}
