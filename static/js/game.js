/* LongWar, by Yves Junqueira.
 *
 *
 */

window.onload = (function() {

	Crafty.init();

	// This will fetch "REPONAME/hexmap-VERSION.js".
	Crafty.modules("../js/thirdparty", {
		hexmap: 'release'
	}, longWarMap)
});

const spriteSize = 64;
const tiles = ["dirt", "mountain", "darkForest", "sand", "forest", "savannah", "water"];
const lastTerrain = 6;


// from:
// http://jquerybyexample.blogspot.com/2012/06/get-url-parameters-using-jquery.html
function GetURLParameter(sParam) {
    var sPageURL = window.location.search.substring(1);
    var sURLVariables = sPageURL.split('&');
    for (var i = 0; i < sURLVariables.length; i++) 
    {
        var sParameterName = sURLVariables[i].split('=');
        if (sParameterName[0] == sParam) 
        {
            return sParameterName[1];
        }
    }
}

function longWarMap() {

	const width = 800;
	const height = 600;

	// 30x30 is able to cover a 800x600 screen.
	const numColumns = 30;
	const numLines = 30;
	// Each individual item in the sprite has spriteSize*spriteSize pixels.
	Crafty.sprite(spriteSize, "../images/hexagons.png", {
		// Sprites from 0 to lastTerrain are used for random terrain generation.
		dirt: [0, 0],
		mountain: [1, 0],
		darkForest: [2, 0],
		sand: [3, 0],
		forest: [4, 0],
		savannah: [5, 0],
		water: [6, 0]
	});

	var hexmap = Crafty.e("Hexmap").hexmap(spriteSize, spriteSize);

	// OH THE HORROR.
	hexmap.loadMap = function(mapWidth, mapHeight, mapFromServer, creationFunc) {
		// mapFromServer is a 30*30 int array.
		this._mapWidth = mapWidth;
		this._mapHeight = mapHeight;
		for (j = 0; j < this._mapHeight; j++) {
			for (i = 0; i < this._mapWidth; i++) {
				t = -1;
				pos = j * (this._mapHeight) + i;
				if ((mapFromServer) && (mapFromServer[pos])) {
					t = mapFromServer[pos]
				}
				if (t == -1) {
					t = 0;
				}
				tile = creationFunc(t);

				tile.addComponent("HexmapNode").hexmapNode(i, j);
				this.placeTile(i, j, tile).setTile(i, j, tile);
			}
		}
	}

	var mapSeed = GetURLParameter("seed");

	$.ajax({
		url: '/createmap?jsoncallback=?&seed=' + mapSeed,
		async: true,
		// Using JSONP even though this will work on the same domain because I
		// needed to use localhost for development. That isn't the case
		// anymore but I'll keep using jsonp anyway just in case.
		dataType: 'jsonp',
		timeout: 3000,

		error: function(xhr, textStatus) {
			// TODO: Show an error on the web page. The error is usually a
			// timeout, because jQuery can't run this callback properly for
			// jsonp.
			console.log("error:", textStatus);
		},
		success: function(data) {

			// Fetch a random map from the server.
			var mapFromServer = [];
			$.each(data, function(i, v) {
				mapFromServer.push(v);
			});

			hexmap.loadMap(numColumns, numLines, mapFromServer, function(tileIndex) {
				var terrain = tiles[tileIndex];
				return Crafty.e("2D, DOM, Mouse, " + terrain).attr({
					tileIndex: tileIndex
				}).bind("Click", function() {
					var nextTile = (this.tileIndex + 1) % (lastTerrain + 1);
					this.removeComponent(tiles[this.tileIndex])
					this.addComponent(tiles[nextTile])
					this.tileIndex = nextTile;
				}).bind("MouseOver", function() {
				/*
					if (grass) {
						if (grass) grass.forEach(function(tile) {
						tile.removeComponent("grass").addComponent("forrest");
					});

				grass = hexmap.findpath(grass, this);
				if (grass) grass.forEach(function(tile) {
					tile.removeComponent("forrest").addComponent("grass");
				});
					}
				*/
				});
			});


		}
	});
};
