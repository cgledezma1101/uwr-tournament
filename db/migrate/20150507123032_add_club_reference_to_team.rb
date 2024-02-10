class AddClubReferenceToTeam < ActiveRecord::Migration[6.0]
	def change
		add_reference :teams, :club, index: true
	end
end
