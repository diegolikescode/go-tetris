import Card from "../Card";

import { Container } from "./styles";

export interface SingleUser {
	id: number;
	name: string;
	city: string;
	country: string;
	favorite_sport: string;
}

const CardContainer = ({ usersData, ...props }: any) => {
	return (
		<Container>
			{usersData.count > 0 &&
				usersData.rows.map((singleUser: SingleUser) => (
					<Card key={singleUser.id} userData={singleUser} />
				))}
		</Container>
	);
};

export default CardContainer;
