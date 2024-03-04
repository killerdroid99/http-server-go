import { useParams } from "@solidjs/router";
import { Response, TPost } from "../types";
import { Show, createResource } from "solid-js";
import Time from "../components/Time";

const fetchPostById = async (postId: string): Promise<Response<TPost>> => {
	const res = await fetch(`http://localhost:8080/post/${postId}`);

	return res.json();
};

function PostPage() {
	const params = useParams();

	const [data] = createResource(params.id, fetchPostById);

	// const createdAt =  as string;

	console.log();

	return (
		<main class="pt-10 p-8 lg:px-[15dvw] grid gap-8 justify-start">
			<Show when={data()?.message === "success"} fallback={<p>loading...</p>}>
				<div>
					<p class="capitalize">{data()?.data.value.authorName}</p>
					<Time time={data()?.data.value.createdAt as string} />
				</div>
				<h1 class="text-6xl font-bold break-words max-w-xl">
					{data()?.data.value.title}
				</h1>
				<article class="text-lg markdown">
					<p innerHTML={data()?.data.value.body}></p>
				</article>
			</Show>
		</main>
	);
}

export default PostPage;
