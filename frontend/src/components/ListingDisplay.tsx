import useListingStore from "../store/useListingStore";

const ListingDisplay = () => {
	const { listings, removeListing } = useListingStore();

	return (
		<div>
			<h2>Listings</h2>
			<ul>
				{listings.map((listing) => (
					<li key={listing.id}>
						<h3>{listing.title}</h3>
						<button onClick={() => removeListing(listing.id)}>
							Remove
						</button>
					</li>
				))}
			</ul>
		</div>
	);
};

export default ListingDisplay;
