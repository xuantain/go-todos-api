document.addEventListener('DOMContentLoaded', function(e) {

    $('.btn-del__todo').click(function(e) {
        console.log('Delete Todo', $(this).data('id'));
        let todoId = $(this).data('id');

        let text;
        if (confirm('Delete Todo ' + todoId + ' ?')) {
            text = 'You pressed OK!';
            let url = '/api/users/1/todos/' + todoId;
            fetch(url, { method: 'DELETE' })
                .then(res => {
                    console.log(res);
                })
                .catch(errorMsg => { console.log(errorMsg); });
        } else {
            text = 'You canceled!';
        }

        console.log(text);
    });
});
