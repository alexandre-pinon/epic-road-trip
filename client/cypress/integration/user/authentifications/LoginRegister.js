describe('The Register and Login page', () => {

    it('Test Login button from the home page', () => {
        cy.visit('http://localhost:3000') // change URL to match your dev URL
        cy.wait(500)
        cy.get('LoginPage').click()
    })

    it('Test Register button from the home page', () => {
        cy.visit('http://localhost:3000') // change URL to match your dev URL
        cy.wait(500)
        cy.get('RegisterPage').click()
    })

})