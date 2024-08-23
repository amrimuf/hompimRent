import { Listing } from "../types/interfaces";
import api from "./api";

export const fetchListings = async (): Promise<Listing[]> => {
	const response = await api.get<Listing[]>("/protected/listings");
	return response.data;
};
