import React from "react";
import { useQuery } from "@tanstack/react-query";
import { fetchListings } from "../services/listingService";
import { Listing } from "../types/interfaces";

const ListingDisplay: React.FC = () => {
	const { data, isLoading, isError, error } = useQuery<Listing[], Error>({
		queryKey: ["listings"],
		queryFn: fetchListings,
		staleTime: 5 * 60 * 1000, // 5 minutes
		gcTime: 30 * 60 * 1000, // 30 minutes (formerly cacheTime)
	});

	if (isLoading) {
		return <div>Loading listings...</div>;
	}

	if (isError) {
		return <div>Error loading listings: {error.message}</div>;
	}

	return (
		<div className="container mx-auto p-4">
			<h1 className="text-2xl font-bold mb-4">Listings</h1>
			{data && data.length > 0 ? (
				<div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-4">
					{data.map((listing: Listing) => (
						<div
							key={listing.id}
							className="border rounded-lg p-4 shadow hover:shadow-lg transition-shadow"
						>
							<h2 className="text-xl font-semibold mb-2">
								{listing.title}
							</h2>
							<p className="text-gray-700 mb-2">
								{listing.description}
							</p>
							<p className="text-green-600 font-bold">
								${listing.price.toFixed(2)}
							</p>
						</div>
					))}
				</div>
			) : (
				<div>No listings available.</div>
			)}
		</div>
	);
};

export default ListingDisplay;
