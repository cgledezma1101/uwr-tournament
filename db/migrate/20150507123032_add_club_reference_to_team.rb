class AddClubReferenceToTeam < ActiveRecord::Migration
	def change
		add_reference :teams, :club, index: true
	end
end
