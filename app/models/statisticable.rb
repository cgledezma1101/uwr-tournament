# This is an "interface" that defines models that can retrieve statistics
class Statisticable < ActiveRecord::Base
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

  # @return [Array<Hash<String, Integer>>] An array where each item is a hash containing the name of a team and the amount of goals it has received
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

  # @return [Array<Hash<String, Integer>>] An array where each item contains the name of a player and the amount of goals it has scored
  def top_scorers
    raise "This method hasn't been implemented"
  end

  # @param [Team] team A team for which it should be determined the amount of games it has won in this set of statistics
  # @return [Integer] The amount of games that the given team has won
  def won_games(team)
    raise "This method hasn't been implemented"
  end
end
