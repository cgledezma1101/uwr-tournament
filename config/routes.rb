Rails.application.routes.draw do
  root 'public#home'

  devise_for :users, controllers: { omniauth_callbacks: "users/omniauth_callbacks" }

  get '/statistics', to: 'public#statistics'

  resources :club_admins, only: [:new, :create, :destroy] do
    member do
      get 'confirm_destroy'
    end
  end

  resources :club_join_requests, only: [:new, :create] do
    member do
      get 'accept'
      get 'decline'
    end
  end

  resources :clubs, only: [:show, :new, :create, :edit, :update]

  resources :game_events, only: [:create, :destroy]

  resources :games, only: [:create, :new, :show, :destroy] do
    member do
      get 'end'
      post 'finalize'
      post 'start'
      post 'add_score'
      post 'remove_score'
    end

    collection do
      get 'add_chronometer'
      get 'external_scoreboard'
      get 'new_auto_leaderboard'
      get 'new_auto_game_result'
    end

  end

  resources :invitations, only: [:new, :create] do
    member do
      post 'accept'
      post 'decline'
    end
  end

  resources :players, only: [:new, :create, :destroy] do
    member do
      get 'confirm_destroy'
    end
  end

  resources :scores, only: [:create]

  resources :stages, only: [:new, :create, :show, :destroy] do
  end

  resources :teams, only: [:new, :create, :show, :destroy, :edit, :update] do
    member do
      get 'confirm_destroy'
      get 'players'
    end
  end

  resources :tournament_admins, only: [:new, :create, :destroy] do
    member do
      get 'confirm_destroy'
    end
  end

  resources :tournament_invitations, only: [:new, :create, :destroy]

  resources :tournament_teams, only: [:new, :create]

  resources :tournaments, only: [:new, :create, :show] do
    member do
      get 'all_games'
    end
  end

  resources :user_clubs, only: [:destroy] do
    member do
      get 'confirm_destroy'
    end
  end

  resources :users, only: [:show]
end
