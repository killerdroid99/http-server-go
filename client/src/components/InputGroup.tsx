import clsx from "clsx";
import { JSX, Show } from "solid-js";
import { TError } from "../types";
import { Warning } from "../icons/Warning";

interface TProps {
	inputAttrs: Omit<
		JSX.InputHTMLAttributes<HTMLInputElement>,
		"class" | "classList" | "style"
	>;
	errors: () => TError | undefined;
	fieldName: string;
}

function InputGroup({ inputAttrs, errors, fieldName }: TProps) {
	return (
		<div class="grid">
			<input
				{...inputAttrs}
				class={clsx(
					"outline-none border-none w-full dark:bg-zinc-800 p-3 rounded-full ring-2  ",
					{
						"placeholder:text-rose-600 text-rose-500 ring-red-500":
							errors()?.field === fieldName,
						"focus-visible:ring-blue-600 ring-slate-700 dark:ring-zinc-700 placeholder:text-slate-600 dark:placeholder:text-zinc-600":
							errors()?.field !== fieldName,
					}
				)}
			/>
			<Show when={errors()?.field === fieldName}>
				<small class="ml-5 text-rose-500 inline-flex items-center gap-1 leading-8">
					<Warning /> {errors()?.value}
				</small>
			</Show>
		</div>
	);
}

export default InputGroup;
