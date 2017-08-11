class AddPasswordToTournamentTeams < ActiveRecord::Migration
  def change
    add_column :tournament_teams, :password, :string
  end
end
