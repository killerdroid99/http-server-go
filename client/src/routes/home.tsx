import { For, useContext } from "solid-js";
import Post from "../components/Post";
import { PostsContext } from "../context/posts.context";
import { A } from "@solidjs/router";

function Home() {
	const Posts = useContext(PostsContext);

	return (
		<main class="pt-10 p-8 lg:px-[15dvw] grid gap-8 justify-center">
			<For each={Posts?.data()?.data.value}>
				{(post, index) => (
					<A href={`/post/${post.id}`} data-index={index()}>
						<Post {...post} />
					</A>
				)}
			</For>
		</main>
	);
}

export default Home;
