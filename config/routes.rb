UwrTournament::Application.routes.draw do
  # You can have the root of your site routed with "root"
  root 'games#new'

  resources :games, only: [:create, :new, :show] do
    resources :scores, only: [:create]
  end

  resources :teams, only: [] do
    get 'players', on: :member
  end
end
