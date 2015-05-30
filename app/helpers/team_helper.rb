module TeamHelper
  # Allows the creation of a method that will render management actions for a team's players
  #
  # @param [Team] team The team for which the actions will be performed
  #
  # @return [Proc] A method that can be called with a particular player and that will render the actions that can be performed on that player for the mentioned team
  def create_player_actions_proc(team)
    Proc.new do |player|
      render 'teams/player_actions', player: player
    end
  end
end
