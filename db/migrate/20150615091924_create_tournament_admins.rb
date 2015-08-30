class CreateTournamentAdmins < ActiveRecord::Migration
  def change
    create_table :tournament_admins do |t|
      t.integer :user_id
      t.integer :tournament_id

      t.timestamps
    end
  end
end
