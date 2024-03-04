export interface Response<T> {
	status: number;
	message: "success" | "error";
	data: {
		value: T;
	};
}

export interface TPost {
	id: string;
	title: string;
	body: string;
	createdAt: string;
	updatedAt: string;
	authorId: string;
	authorName: string;
}

export interface TError {
	field: string;
	value: string;
}

export interface TAuthData {
	userId: string;
	userName: string;
}
