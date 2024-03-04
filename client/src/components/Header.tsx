import { Match, Switch, useContext } from "solid-js";
import { AuthContext } from "../context/auth.context";
import { A } from "@solidjs/router";
import { Add } from "../icons/Add";

function Header() {
	const Auth = useContext(AuthContext);

	const handleLogout = async () => {
		const res = await fetch("http://localhost:8080/logout", {
			method: "POST",
			credentials: "include",
		});

		Auth?.refetch();
		console.log(await res.json());
	};

	return (
		<header class="bg-gradient-to-b from-white to-white/30 dark:from-zinc-900 dark:to-zinc-900/30 z-10 backdrop-blur-sm p-4 flex items-center justify-between lg:px-[13dvw] sticky top-0">
			<A href="/" class="text-xl font-bold">
				Client
			</A>

			<nav class="flex items-center gap-2">
				<Switch fallback={<p>Unknown error</p>}>
					<Match when={Auth?.data()?.message === "success"}>
						<span class="outline-none border-none p-2 rounded focus-visible:ring-4 ring-rose-400/80 font-bold">
							{Auth?.data()?.data.value.userName}
						</span>
						<A
							href="/create-post"
							class="outline-none border-none hover:bg-slate-400/20 dark:hover:bg-slate-500/10 p-2 rounded focus-visible:ring-4 ring-slate-400/80 text-2xl"
							title="Add post"
						>
							<Add />
						</A>
						<button
							class="outline-none border-none bg-rose-500 text-rose-50 p-2 rounded focus-visible:ring-4 ring-rose-400/80 font-bold"
							onclick={handleLogout}
						>
							Logout
						</button>
					</Match>
					<Match when={Auth?.data()?.message === "error"}>
						<A
							href="/login"
							class="outline-none border-none hover:bg-slate-400/20 dark:hover:bg-slate-500/10 p-2 rounded focus-visible:ring-4 ring-slate-400/80 font-bold"
							activeClass="text-indigo-500"
						>
							Login
						</A>
						<A
							href="/register"
							class="outline-none border-none hover:bg-slate-400/20 dark:hover:bg-slate-500/10 p-2 rounded focus-visible:ring-4 ring-slate-400/80 font-bold"
							activeClass="text-indigo-500"
						>
							Register
						</A>
					</Match>
				</Switch>
			</nav>
		</header>
	);
}

export default Header;
