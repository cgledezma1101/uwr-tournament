class ClubAdmin < ActiveRecord::Base
  belongs_to :user
  belongs_to :club

  validate :user, uniqueness: { scope: :club }

  # Returns the full name of the administrator associated by this relationship
  #
  # @return [String] The name of the associated user
  def admin_name
    self.user.name
  end
end
