type GetData = {
	count: number;
	rows: Array<object>;
};

export interface CustomError {
	msg: string;
	isError: boolean;
}

export const getData = async (q?: string): Promise<GetData | CustomError> => {
	try {
		const res = await fetch(
			`http://localhost:3069/api/users?${q ? "q=" + q : ""}`
		);

		console.log(res.status);

		if (res.status === 204)
			return {
				msg: "what? I didn't find anything",
				isError: true,
			};

		return await res.json();
	} catch (err) {
		return {
			msg: "The search is dead. Long live the search!",
			isError: true,
		};
	}
};

export const postData = async (file: any): Promise<string | CustomError> => {
	try {
		const data = new FormData();
		data.append("sweetCsv", file);
		await fetch("http://localhost:3069/api/files", {
			body: data,
			method: "POST",
		});

		return "OK";
	} catch (err) {
		console.log("err postData", err);
		return { msg: "SHIEESH, YOUR UPLOAD DIDN'T MAKE IT.", isError: true };
	}
};
