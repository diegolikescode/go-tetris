import {Container, Strong, P} from './styles'

const Card = ({
	userData: { name, city, country, favorite_sport },
	...props
}: any) => {
	return (
		<Container>
			<strong>Name: {name}</strong>
			<p>City: {city}</p>

			<p>Country: {country}</p>
			<p>Favorite Sport: {favorite_sport}</p>
		</Container>
	);
};

export default Card;
