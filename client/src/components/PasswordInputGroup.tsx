import clsx from "clsx";
import { JSX, Show } from "solid-js";
import { TError } from "../types";
import { Warning } from "../icons/Warning";
import { EyeClosedOutline } from "../icons/EyeClosedOutline";
import { EyeOutline } from "../icons/EyeOutline";

interface TProps {
	inputAttrs: Omit<
		JSX.InputHTMLAttributes<HTMLInputElement>,
		"class" | "classList" | "style" | "type"
	>;
	errors: () => TError | undefined;
	isPasswordVisible: () => boolean;
	setIsPasswordVisible: (value: boolean) => boolean;
}

function PasswordInputGroup({
	inputAttrs,
	errors,
	isPasswordVisible,
	setIsPasswordVisible,
}: TProps) {
	return (
		<div class="grid relative">
			<input
				{...inputAttrs}
				type={isPasswordVisible() ? "text" : "password"}
				class={clsx(
					"outline-none border-none w-full dark:bg-zinc-800 p-3 rounded-full ring-2  ",
					{
						"placeholder:text-rose-600 text-rose-500 ring-red-500":
							errors()?.field.includes("word"),
						"focus-visible:ring-blue-600 ring-slate-700 dark:ring-zinc-700 placeholder:text-slate-600 dark:placeholder:text-zinc-600":
							!errors()?.field.includes("word"),
					}
				)}
			/>
			<Show when={errors()?.field.includes("word")}>
				<small class="ml-5 text-rose-500 inline-flex items-center gap-1 leading-8">
					<Warning /> {errors()?.value}
				</small>
			</Show>

			<div
				class="absolute top-3 right-3 cursor-pointer text-2xl text-slate-700 dark:text-zinc-500"
				onclick={() => setIsPasswordVisible(!isPasswordVisible())}
			>
				<Show when={isPasswordVisible()}>
					<EyeClosedOutline />
				</Show>
				<Show when={!isPasswordVisible()}>
					<EyeOutline />
				</Show>
			</div>
		</div>
	);
}

export default PasswordInputGroup;
