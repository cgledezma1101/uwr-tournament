class Invitation < ActiveRecord::Base
  belongs_to :club
  belongs_to :user

  validates :club, presence: true
  validates :user, presence: true

  validate :one_invitation_per_club

  private

  # Validates the uniqueness of the user in the scope of the club
  def one_invitation_per_club
    if !user.nil?
      duplicate = Invitation.find_by(club: self.club, user: self.user)
      
      if !duplicate.nil? && duplicate.id != self.id
        self.errors.add(:user, I18n.t('invitations.only_one_per_club'))
      end
    end
  end
end
