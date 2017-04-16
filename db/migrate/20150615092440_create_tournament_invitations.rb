class CreateTournamentInvitations < ActiveRecord::Migration
	def change
		create_table :tournament_invitations do |t|
			t.integer :club_id
			t.integer :tournament_id

			t.timestamps
		end
	end
end
