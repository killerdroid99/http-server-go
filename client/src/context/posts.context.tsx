import { createContext, createResource, ParentProps, Resource } from "solid-js";
import { Response, TPost } from "../types";

const fetchAllPosts = async (): Promise<Response<TPost[]>> => {
	const res = await fetch("http://localhost:8080/posts");

	return res.json();
};

export const PostsContext = createContext<{
	data: Resource<Response<TPost[]>>;
	refetch: (
		info?: unknown
	) =>
		| Response<TPost[]>
		| Promise<Response<TPost[]> | undefined>
		| null
		| undefined;
}>();

export function PostsContextProvider(props: ParentProps) {
	const [data, { refetch }] = createResource(fetchAllPosts);

	return (
		<PostsContext.Provider
			value={{
				data,
				refetch,
			}}
		>
			{props.children}
		</PostsContext.Provider>
	);
}
