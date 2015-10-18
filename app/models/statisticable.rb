# This is an "interface" that defines models that can retrieve statistics
class Statisticable
  # @param [Team] The team for whom the amount of goals is desired
  # @return [Integer] A numeric representation of the balance between the goals the team has scored and received
  def goals_average(team)
    raise "This method hasn't been implemented"
  end

  # @param [Team] The team for whom the amount of goals is desired
  # @return [Integer] The amount of goals the team has received in this statistics set
  def goals_received(team)
    raise "This method hasn't been implemented"
  end

  # @param [Team] The team for whom the amount of goals is desired
  # @return [Integer] The amount of goals the team has scored in this statistics set
  def goals_scored(team)
    raise "This method hasn't been implemented"
  end

  # @return [Array<Array<Team, Integer>>] An array containing tuples whose first item is a team, and whose second item is the amount of goals that team has received. The array should be sorted in ascending order by goals received
  def least_defeated
    raise "This method hasn't been implemented"
  end

  # @param [Team] team A team for which it should be determined the amount of games it has lost in this set of statistics
  # @return [Integer] The amount of games that the given team has lost
  def lost_games(team)
    raise "This method hasn't been implemented"
  end

  # @return [Array<Team>] Determines the teams that participate in this set of statistics
  def participating_teams
    raise "This method hasn't been implemented"
  end

  # @param [Team] team The team for whom the amount of points is queried
  # @return [Integer] The amount of points this team has accumulated in this statistics set
  def points_for(team)
    raise "This method hasn't been implemented"
  end

  # @param [Team] team A team for which it should be determined the amount of games it has tied in this set of statistics
  # @return [Integer] The amount of games that the given team has tied
  def tied_games(team)
    raise "This method hasn't been implemented"
  end

  # @return [Array<Array<Player, Integer>>] An array where each element contains a tuple whose first item is a player, and second item is the amount of goals that player has scored in the tournament. The array should be sorted in descending order by goals scored
  def top_scorers
    raise "This method hasn't been implemented"
  end

  # @param [Team] team A team for which it should be determined the amount of games it has won in this set of statistics
  # @return [Integer] The amount of games that the given team has won
  def won_games(team)
    raise "This method hasn't been implemented"
  end
end
