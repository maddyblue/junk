function SchalmeiCtrl($scope, $http) {
	$scope.show = function(name, val) {
		$scope.shown = name;

		if (name == 'show-ranks') {
			$http.get($('#show-rank-list').attr('data-url')).
				success(function(data, status) {
					$scope.ranks = data;
				});
		}
		else if (name == 'get-rank') {
			$http({
				method: 'GET',
				url: val.Url
			}).success(function(data, status) {
				$scope.rank = data;
			});
		}
	};

	$scope.isActive = function(name) {
		return $scope.shown == name ? 'active' : '';
	};

	$scope.createRank = function() {
		if (!$scope.createRankName)
			return;

		$http.post($('#create-rank-form').attr('action'), $scope.createRankName).
			success(function(data, status) {
				$scope.show('get-rank', data);
			});
	};

	$scope.uploadNote = function() {
		var form = $('#file_form');
		$scope.uploadImageUrl($scope.rank.Url, function(data) {
			var url = JSON.parse(data);
			form.attr('action', url);
			form.ajaxSubmit({
				success: function(data) {
					$scope.rank = JSON.parse(data);
					$scope.$digest();
				},
				error: function(data) {
					alert("upload error");
				}
			});
		});
	};

	$scope.uploadImageUrl = function(url, callback) {
		$.ajax({
			url: url
		}).done(callback);
	};
}

angular.module('schalmei', [])
	.directive('schGraph', function() {
		return function(scope, element, attrs) {
			var url;

			function updateGraph() {
				element.attr('id', '_' + Math.random()).empty();
				var options = {
					'width': 800,
					'height': 500,
					curveType: "function"
				};
				var wrapper = new google.visualization.ChartWrapper({
					chartType: 'LineChart',
					options: options,
					dataSourceUrl: url,
					containerId: element.attr('id')
				});
				wrapper.draw();
			}

			scope.$watch(attrs.schGraph, function(value) {
				url = value;
				updateGraph();
			});
		}
	});
