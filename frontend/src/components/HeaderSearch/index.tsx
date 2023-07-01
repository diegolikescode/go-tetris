import { ChangeEvent, useRef, useState } from "react";
import {
	Container,
	SearchBar,
	ButtonSubmitSearch,
	InputCsv,
	SubContainer,
	UploadMsg
} from "./styles";
import { postData } from "../../services/api";

const HeaderSearch = ({ handleSearch }: any) => {
	const searchRef = useRef<HTMLInputElement>(null);
	const [file, setFile] = useState<File>();
	const [uploadMsg, setUploadMsg] = useState<string>("");

	const handleUploadFile = async () => {
		if (!file) return;

		console.log(file);
		const postedCsv = await postData(file);

		if(typeof postedCsv === "string") {

			setUploadMsg("HEY, NICE UPLOAD YOU GOT THERE")
			setTimeout(() => {
				setUploadMsg("")
			}, 5000)
		} else {

			setUploadMsg(postedCsv.msg)
			setTimeout(() => {
				setUploadMsg("")
			}, 5000)

		}
	};

	const onEnterSearch = (ev: any) => {
		if (searchRef.current) {
			if (ev.key === "Enter") {
				handleSearch(searchRef.current.value || "");
			}
		}
	};
	return (
		<Container>
			<SubContainer>
				<InputCsv
					type="file"
					accept=".csv"
					onChange={(e: ChangeEvent<HTMLInputElement>) =>
						setFile(
							e.target && e.target.files !== null
								? e.target.files[0]
								: undefined
						)
					}
				/>
				<ButtonSubmitSearch onClick={() => handleUploadFile()}>
					UPLOAD FILE
				</ButtonSubmitSearch> <UploadMsg>{uploadMsg}</UploadMsg>

			</SubContainer>
			<SubContainer>
				<SearchBar
					type="search"
					ref={searchRef}
					onKeyUp={(e: any) => onEnterSearch(e)}
				/>
				<ButtonSubmitSearch
					onClick={() => {
						handleSearch(
							searchRef.current ? searchRef.current.value : ""
						);
					}}
				>
					SEARCH
				</ButtonSubmitSearch>
			</SubContainer>
		</Container>
	);
};

export default HeaderSearch;
