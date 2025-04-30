document.addEventListener("DOMContentLoaded", function(e) {

    $('.btn-del__todo').click(function(e) {
        console.log('Delete Todo', $(this).data('id'));
        var todoId = $(this).data('id');
        var url = '/api/users/1/todos/' + todoId;
        fetch(url, { method: 'DELETE' })
            .then(Result => Result.json())
            .then(string => {

                // Printing our response 
                console.log(string);

                // Printing our field of our response
                console.log(`Title of our response :  ${string.title}`);
            })
            .catch(errorMsg => { console.log(errorMsg); }); 
        // $.ajax({
        //     type: 'DELETE',
        //     url : url,
        //     success: function(res) {
        //         console.log(res);
        //     },
        //     error: function (e) {
        //         console.error(e);
        //     }
        // })
        // .always(function (data) {
        //     console.log('all done');
        // });
    });
});
