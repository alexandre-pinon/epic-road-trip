describe('Auth', () => {

  it('Login', () => {
    cy.visit('http://localhost:3000/login')

    cy.get('[data-attr=login-email]').type('hansohee@gmail.com').should('have.value', 'hansohee@gmail.com').blur()
    cy.get('[data-attr=login-email]', { timeout: 5000 }).should('be.visible')

    cy.get('[data-attr=login-password]').type('password').should('have.value', 'password')
    cy.get('[data-attr=login-confirm]').click()

    cy.location('pathname').should('eq', '/')
  })

  it('Redirect to appropriate place after register', () => {
    cy.visit('http://localhost:3000/register')
    cy.location('pathname').should('include', '/register')

    cy.wait(1000)

    cy.get('[data-attr=register-confirm]').click()
    cy.location('pathname').should('include', '/login')

    cy.wait(1000)

    cy.get('[data-attr=login-confirm]').click()
    cy.location('pathname').should('eq', '/')
  })
})