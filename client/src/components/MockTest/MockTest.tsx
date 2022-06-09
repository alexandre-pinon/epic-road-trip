import React, { useState, useCallback } from 'react'

export const MockTest = () => {
  const [username, setUsername] = useState('')

  const [userData, setUserData] = useState(null)

  const handleUsernameChange = useCallback((event: { target: { value: React.SetStateAction<string> } }) => {
    setUsername(event.target.value)
  }, [])

  const handleFormSubmit = useCallback(
    (event: { preventDefault: () => void }) => {
      event.preventDefault()

      fetch('/login', {
        method: 'POST',
        body: JSON.stringify({
          username,
        }),
      })
        .then((res) => res.json())
        .then(setUserData)
    },
    [username]
  )

  if (userData) {
    return (
      <div>
        <h1>
          <span data-testid="firstName">{userData['firstName']}</span>{' '}
          <span data-testid="lastName">{userData['lastName']}</span>
        </h1>
        <p data-testid="userId">{userData['id']}</p>
      </div>
    )
  }

  return (
    <form onSubmit={handleFormSubmit}>
      <div>
        <label htmlFor="username">Username:</label>
        <input
          id="username"
          name="username"
          value={username}
          onChange={handleUsernameChange}
        />
        <button type="submit">Submit</button>
      </div>
    </form>
  )
}