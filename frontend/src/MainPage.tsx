import { useEffect, useState } from "react";
import CardContainer from "./components/CardContainer";
import HeaderSearch from "./components/HeaderSearch";
import { CustomError, getData } from "./services/api";

const MainPage = () => {
	const [usersData, setUsersData] = useState({});
	const [error, setError] = useState<CustomError>({
		msg: "dang it! I was wrong ;-;",
		isError: false,
	});
	const handleSearch = (searchRef: string) => {
		const searchValue = searchRef || "";
		getData(searchValue).then((res: any) => {
			if (res.isError) {
				setError({ msg: res.msg, isError: true });
			} else {
				setUsersData(res);
				setError({msg: "", isError: false})
			}
		});
	};

	useEffect(() => {
		handleSearch("");
	}, []);

	return (
		<>
			<HeaderSearch handleSearch={handleSearch} />
			{!error.isError ? <CardContainer usersData={usersData} /> : <h1>{error.msg}</h1>}
		</>
	);
};

export default MainPage;
