import { Show } from "solid-js";
import { TPost } from "../types";
import Time from "./Time";

function Post({ title, body, createdAt, updatedAt, authorName }: TPost) {
	const rtf = new Intl.DateTimeFormat("en-us", {
		dateStyle: "medium",
	});
	return (
		<div class="grid ring-slate-700 dark:ring-zinc-700 ring-2 rounded-3xl p-6 gap-2 md:w-[90dvw] xl:max-w-4xl">
			<div class="grid">
				<div class="font-medium text-sm capitalize leading-5">{authorName}</div>
				<span class="text-xs font-medium dark:text-zinc-500 text-slate-500">
					Posted on <Time time={createdAt as string} />{" "}
					<Show when={createdAt !== updatedAt}>
						<span>(edited)</span>
					</Show>
				</span>
			</div>
			<div class="grid">
				<h3 class="text-2xl font-bold text-slate-800 dark:text-slate-50">
					{title}
				</h3>
			</div>
			<p class="text-slate-900 dark:text-slate-100">
				{body.length > 50 ? body.slice(0, 50) + "..." : body}
			</p>
		</div>
	);
}

export default Post;
