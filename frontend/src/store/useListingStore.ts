import { create } from "zustand";
import { ListingState } from "../types/interfaces";

const useListingStore = create<ListingState>((set) => ({
	listings: [],
	addListing: (listing) =>
		set((state) => ({
			listings: [...state.listings, listing],
		})),
	removeListing: (id) =>
		set((state) => ({
			listings: state.listings.filter((listing) => listing.id != id),
		})),
	updateListing: (id, updatedListing) =>
		set((state) => ({
			listings: state.listings.map((listing) =>
				listing.id === id ? { ...listing, ...updatedListing } : listing
			),
		})),
}));

export default useListingStore;
