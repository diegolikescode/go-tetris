import styled from "styled-components";

export const Container = styled.div`
	display: flex;
	align-items: center;
	justify-content: space-around;
	background-color: orange;
	width: 100%;
	max-height: 280px;
	gap: 12px;
	padding: 10px 12px;
`;

export const SubContainer = styled.div`
	display: flex;
	flex-direction: column;
	gap: 8px;
	color: black;
`;

export const SearchBar = styled.input`
	height: 24px;
	width: 100%;
	color: black;
	padding: 6px 0px 6px 4px;
`;

export const ButtonSubmitSearch = styled.button`
	color: black;
	padding: 2px 4px;
	height: 24px;
	width: 120px;
`;

export const InputCsv = styled.input``;

export const UploadMsg = styled.p`
	color: black;
`;
