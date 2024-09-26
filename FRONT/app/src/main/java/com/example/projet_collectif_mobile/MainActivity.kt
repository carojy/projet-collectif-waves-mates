package com.example.projet_collectif_mobile

import android.content.Intent
import android.os.Bundle
import android.view.LayoutInflater
import android.view.View
import android.view.ViewGroup
import android.widget.ArrayAdapter
import android.widget.Button
import android.widget.ListView
import android.widget.TextView
import androidx.activity.enableEdgeToEdge
import androidx.activity.viewModels
import androidx.appcompat.app.AppCompatActivity
import androidx.lifecycle.Observer
import com.example.projet_collectif_mobile.models.SurfSpot
import com.example.projet_collectif_mobile.ui.screens.SpotsViewModel

class MainActivity : AppCompatActivity() {
    // Use the viewModels method to instantiate the ViewModel
    private val spotsViewModel: SpotsViewModel by viewModels()

    override fun onCreate(savedInstanceState: Bundle?) {
        super.onCreate(savedInstanceState)
        enableEdgeToEdge()
        setContentView(R.layout.activity_main)

        val listView=findViewById<ListView>(R.id.list_view)

        // Observe changes of spotsUiState in the ViewModel
        // spotsList contains the list of spots, to be used in the ArrayAdapter
        spotsViewModel.spotsUiState.observe(this, Observer { spotsList ->

            val arrayAdapter: ArrayAdapter<SurfSpot> = object : ArrayAdapter<SurfSpot>(
                this, R.layout.list_item, spotsList)
            {
                override fun getView(position: Int, convertView: View?, parent: ViewGroup): View {
                    // Reuse converted views or create new ones if necessary
                    val inflater = LayoutInflater.from(context)
                    val view = convertView ?: inflater.inflate(R.layout.list_item, parent, false)

                    // Retrieve Spot object for current position
                    val spot = getItem(position)

                    val spotNameTextView = view.findViewById<TextView>(R.id.spot_name)
                    val spotLocationTextView = view.findViewById<TextView>(R.id.spot_location)

                    spotNameTextView.text = spot?.surfBreak?.joinToString(", ")
                    spotLocationTextView.text = spot?.address


                    val button = view.findViewById<Button>(R.id.spot_btn)
                    button.setOnClickListener {
                        //Create an Intent to launch SpotDetailActivity
                        val intent = Intent(applicationContext, SpotActivity::class.java)

                        // Add Spot details to Intent
                        intent.putExtra("SPOT_NAME", spot?.surfBreak?.joinToString(", "))
                        intent.putExtra("SPOT_LOCATION", spot?.address)

                        startActivity(intent)
                    }
                    return view
                }
            }
            listView.adapter = arrayAdapter
        })
        spotsViewModel.getSpotsList()
    }
}