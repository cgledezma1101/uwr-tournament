class AddPasswordToTournamentTeams < ActiveRecord::Migration[6.0]
  def change
    add_column :tournament_teams, :password, :string
  end
end
