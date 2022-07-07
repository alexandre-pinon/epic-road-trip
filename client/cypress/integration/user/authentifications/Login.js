describe('Login page', () => {

    it('check text on the Login page', () => {
        cy.visit('http://localhost:3000/login')

        cy.get('[data-testid="welcome"]').contains('Welcome 👋!')

        cy.get('[data-testid="registerPage"]')
            .contains('Do not have an account yet? Create account')

        cy.get('[data-testid="signup"]').contains('Sign in')
    })



    it('Check attributes Login page', () => {
        cy.visit('http://localhost:3000/login')
        //Input
        cy.get('[data-testid="input"]')
            .should('have.attr', 'required')

        cy.get('[data-testid="input"]')
            .should('have.attr', 'placeholder')
            .and('equal', 'Your email')

        //Password
        cy.get('[data-testid="password"]')
            .should('have.attr', 'required')

        cy.get('[data-testid="password"]')
            .should('have.attr', 'placeholder')
            .and('equal', 'Your password')
    })



    it('check Register button Page from Login page', () => {
        cy.visit('http://localhost:3000/login')

        cy.get('[data-testid="registerPage"] > .mantine-Text-root').click()
        cy.wait(500)
        cy.location('href').should('include', '/register#')
    })



    it('check input Login page', () => {
        cy.visit('http://localhost:3000/login')

        cy.get('[data-testid="input"]').type('tharick@gmail.com')
        cy.get('[data-testid="input"]').should('have.value', 'tharick@gmail.com')

        cy.get('[data-testid="password"]').type('12345678')
        cy.get('[data-testid="password"]').should('have.value', '12345678')
    })



    it('check Login Valid', () => {
        cy.visit('http://localhost:3000/login')

        cy.get('[data-testid="input"]').type('tharick@gmail.com')
        cy.get('[data-testid="password"]').type('12345678')
        cy.wait(500)
        cy.get('[data-testid="signup"]').click()
        cy.location('href').should('include', '/')

    })
})