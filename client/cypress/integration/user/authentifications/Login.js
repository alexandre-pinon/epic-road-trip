describe('Login page', () => {

    it('check text on the Login page', () => {
        cy.visit('http://localhost:3000/login')

        cy.get('[data-testid="welcome"]').contains('Welcome 👋!')

        cy.get('[data-testid="registerPage"]')
            .contains('Do not have an account yet? Create account')


    })


})