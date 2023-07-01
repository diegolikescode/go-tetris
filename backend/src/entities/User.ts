import { DataTypes } from "sequelize";
import { database } from "../config.js";

const User = database.define("User", {
	id: {
		type: DataTypes.INTEGER,
		autoIncrement: true,
		allowNull: false,
		primaryKey: true,
	},
	name: { type: DataTypes.STRING },
	city: { type: DataTypes.STRING },
	country: { type: DataTypes.STRING },
	favorite_sport: { type: DataTypes.STRING },
});

export default User;
