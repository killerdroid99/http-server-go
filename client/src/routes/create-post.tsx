import clsx from "clsx";
import { Show, createSignal, useContext } from "solid-js";
import { Response } from "../types";
import { useNavigate } from "@solidjs/router";
import { PostsContext } from "../context/posts.context";

function CreatePost() {
	let formData = {
		title: "",
		body: "",
	};
	const Posts = useContext(PostsContext);

	const navigate = useNavigate();

	const [errors, setErrors] = createSignal<{ field: string; value: string }>();

	const handleSubmit = async (e: Event) => {
		e.preventDefault();

		const res: Response<{ field: string; value: string } | undefined> = await (
			await fetch("http://localhost:8080/post", {
				method: "POST",
				body: JSON.stringify(formData),
				credentials: "include",
			})
		).json();

		if (res.message === "error") {
			setErrors(res.data.value);
		}

		navigate("/", { replace: true });
		Posts?.refetch();
	};

	return (
		<main class="pt-10 p-8 lg:px-[15dvw]">
			<form
				class="grid gap-4
      "
				onsubmit={handleSubmit}
			>
				<div>
					<input
						type="text"
						oninput={(e) => {
							setErrors(undefined);
							formData.title = e.target.value;
						}}
						placeholder="Post title"
						class={clsx(
							"outline-none border-none w-full bg-slate-200 dark:bg-zinc-800 p-5 rounded-3xl font-bold text-4xl",
							{
								"placeholder:text-rose-600 text-rose-500":
									errors()?.field === "title",
								"placeholder:text-slate-600 dark:placeholder:text-zinc-600":
									errors()?.field !== "title",
							}
						)}
					/>
					<Show when={errors()?.field === "title"}>
						<small class="ml-4 text-rose-500">{errors()?.value}</small>
					</Show>
				</div>
				<div>
					<textarea
						oninput={(e) => {
							setErrors(undefined);
							formData.body = e.target.value;
						}}
						class={clsx(
							"outline-none border-none w-full bg-slate-100 dark:bg-zinc-800 p-5 rounded-3xl text-xl min-h-80",
							{
								"placeholder:text-rose-600 text-rose-500":
									errors()?.field === "body",
								"placeholder:text-slate-600 dark:placeholder:text-zinc-600":
									errors()?.field !== "body",
							}
						)}
						placeholder="Enter you post content here..."
					/>
					<Show when={errors()?.field === "body"}>
						<small class="ml-4 text-rose-500">{errors()?.value}</small>
					</Show>
				</div>
				<button class="outline-none grid place-items-center border-none w-fit dark:bg-indigo-700 bg-slate-900 text-white p-5 rounded-3xl focus-visible:ring-4 ring-blue-400/80 font-bold ">
					Submit
				</button>
			</form>
		</main>
	);
}

export default CreatePost;
