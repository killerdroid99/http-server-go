function Time({ time }: { time: string }) {
	const rtf = new Intl.DateTimeFormat("en-us", {
		dateStyle: "medium",
	});
	return <time>{rtf.format(new Date(time))}</time>;
}

export default Time;
