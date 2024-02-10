class CreateTournamentAdmins < ActiveRecord::Migration[6.0]
	def change
		create_table :tournament_admins do |t|
			t.integer :user_id
			t.integer :tournament_id

			t.timestamps
		end
	end
end
